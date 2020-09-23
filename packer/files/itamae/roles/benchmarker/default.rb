node.reverse_merge!(
  benchmarker: {
    enable: true,
  },
)

include_cookbook 'benchmarker'

if node[:benchmarker][:enable]
  service 'isuxportal-supervisor.service' do
    action :enable
  end
end