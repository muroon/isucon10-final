# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/services/bench/foo.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("xsuportal/services/bench/foo.proto", :syntax => :proto3) do
    add_message "xsuportal.proto.services.bench.FooRequest" do
      optional :msg, :string, 1
    end
    add_message "xsuportal.proto.services.bench.FooResponse" do
      optional :msg, :string, 1
    end
  end
end

module Xsuportal
  module Proto
    module Services
      module Bench
        FooRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.bench.FooRequest").msgclass
        FooResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.bench.FooResponse").msgclass
      end
    end
  end
end
