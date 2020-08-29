#!/usr/bin/env ruby

$: << File.expand_path('../lib', __dir__)
require 'logger'
require 'griffin'
require 'grpc'
require 'griffin/interceptors/server/logging_interceptor'
require_relative '../grpc/benchmark_queue_service'
require_relative '../grpc/benchmark_report_service'
require 'database'

Xsuportal::Database.connection

Griffin::Server.configure do |c|
  c.bind '0.0.0.0'

  c.services [
    BenchmarkQueueService,
    BenchmarkReportService,
  ]

  c.interceptors [
    Griffin::Interceptors::Server::LoggingInterceptor.new,
  ]

  c.workers 2
  c.pool_size 10, 10
  c.connection_size 10, 10

  c.logger Logger.new('benchmark_server.log')
end

Griffin::Server.run
