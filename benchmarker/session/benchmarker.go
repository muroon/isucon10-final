package session

import (
	"context"
	"errors"
	"fmt"

	"github.com/isucon/isucon10-final/benchmarker/model"
	"github.com/isucon/isucon10-final/benchmarker/proto/xsuportal/resources"
	"github.com/isucon/isucon10-final/benchmarker/proto/xsuportal/services/bench"
	"github.com/isucon/isucon10-final/benchmarker/random"
	"google.golang.org/grpc"
)

type Benchmarker struct {
	queueClinet    bench.BenchmarkQueueClient
	reportClient   bench.BenchmarkReportClient
	scoreGenerator *random.ScoreGenerator

	Team *model.Team
}

func NewBenchmarker(team *model.Team, host string, port int64) (*Benchmarker, error) {
	host = fmt.Sprintf("%s:%d", host, port)

	queueConn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	reportConn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Benchmarker{
		queueClinet:    bench.NewBenchmarkQueueClient(queueConn),
		reportClient:   bench.NewBenchmarkReportClient(reportConn),
		scoreGenerator: random.NewScoreGenerator(),
		Team:           team,
	}, nil
}

func (b *Benchmarker) Do(ctx context.Context, idx int64, team *model.Team) (*bench.ReportBenchmarkResultRequest, error) {
	defer func() {
		// 高負荷時だと Transport が SEGV するので一旦 recover でやりすごすけどどうしようかなこれ
		if err := recover(); err != nil {
			fmt.Printf("%+v\n", err)
		}
	}()

	scoreG := b.scoreGenerator
	req := &bench.ReceiveBenchmarkJobRequest{}
	if team != nil {
		req.TeamId = team.ID
		scoreG = team.ScoreGenerator
	}
	jobResponse, err := b.queueClinet.ReceiveBenchmarkJob(ctx, req)

	score := scoreG.Generate(idx)

	if err != nil {
		return nil, err
	}

	jobHandle := jobResponse.GetJobHandle()
	if jobHandle == nil {
		return nil, nil
	}

	reporter, err := b.reportClient.ReportBenchmarkResult(ctx)
	if err != nil {
		return nil, err
	}

	result := &bench.ReportBenchmarkResultRequest{
		JobId:  jobHandle.GetJobId(),
		Handle: jobHandle.GetHandle(),
		Result: &resources.BenchmarkResult{
			Finished: false,
			Score:    score.Int() / 2,
			ScoreBreakdown: &resources.BenchmarkResult_ScoreBreakdown{
				Raw:       score.Base / 2,
				Deduction: score.Deduction / 2,
			},
			Reason: "",
		},
		Nonce: 1,
	}
	err = reporter.Send(result)
	if err != nil {
		return nil, err
	}
	reportResponse, err := reporter.Recv()
	if err != nil {
		return nil, err
	}
	if reportResponse.AckedNonce != 1 {
		return nil, errors.New("Invalid Nonce")
	}

	result = &bench.ReportBenchmarkResultRequest{
		JobId:  jobHandle.GetJobId(),
		Handle: jobHandle.GetHandle(),
		Result: &resources.BenchmarkResult{
			Finished: true,
			Passed:   !(score.FastFail || score.SlowFail),
			Score:    score.Int(),
			ScoreBreakdown: &resources.BenchmarkResult_ScoreBreakdown{
				Raw:       score.BaseInt(),
				Deduction: score.DeductionInt(),
			},
			Reason: "",
		},
		Nonce: 2,
	}
	err = reporter.Send(result)
	if err != nil {
		return nil, err
	}
	err = reporter.CloseSend()
	if err != nil {
		return nil, err
	}

	reportResponse, err = reporter.Recv()
	if err != nil {
		return nil, err
	}
	if reportResponse.AckedNonce != 2 {
		return nil, errors.New("Invalid Nonce")
	}

	return result, nil
}
