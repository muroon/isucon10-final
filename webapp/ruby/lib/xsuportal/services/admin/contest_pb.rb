# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/services/admin/contest.proto

require 'google/protobuf'

require 'xsuportal/resources/contest_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("xsuportal/services/admin/contest.proto", :syntax => :proto3) do
    add_message "xsuportal.proto.services.admin.UpdateContestRequest" do
      optional :contest, :message, 1, "xsuportal.proto.resources.Contest"
    end
    add_message "xsuportal.proto.services.admin.UpdateContestResponse" do
    end
  end
end

module Xsuportal
  module Proto
    module Services
      module Admin
        UpdateContestRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.admin.UpdateContestRequest").msgclass
        UpdateContestResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.admin.UpdateContestResponse").msgclass
      end
    end
  end
end