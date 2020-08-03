# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/services/contestant/clarifications.proto

require 'google/protobuf'

require 'xsuportal/resources/clarification_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("xsuportal/services/contestant/clarifications.proto", :syntax => :proto3) do
    add_message "xsuportal.proto.services.contestant.ListClarificationsRequest" do
    end
    add_message "xsuportal.proto.services.contestant.ListClarificationsResponse" do
      repeated :clarifications, :message, 1, "xsuportal.proto.resources.Clarification"
    end
    add_message "xsuportal.proto.services.contestant.RequestClarificationRequest" do
      optional :question, :string, 1
    end
    add_message "xsuportal.proto.services.contestant.RequestClarificationResponse" do
      optional :clarification, :message, 1, "xsuportal.proto.resources.Clarification"
    end
  end
end

module Xsuportal
  module Proto
    module Services
      module Contestant
        ListClarificationsRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.contestant.ListClarificationsRequest").msgclass
        ListClarificationsResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.contestant.ListClarificationsResponse").msgclass
        RequestClarificationRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.contestant.RequestClarificationRequest").msgclass
        RequestClarificationResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.contestant.RequestClarificationResponse").msgclass
      end
    end
  end
end