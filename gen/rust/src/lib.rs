// tonic does not derive `Eq` for the gRPC message types, which causes a warning from Clippy. The
// current suggestion is to explicitly allow the lint in the module that imports the protos.
// Read more: https://github.com/hyperium/tonic/issues/1056
#![allow(clippy::derive_partial_eq_without_eq)]

pub mod kintai {
    pub mod v1 {
        tonic::include_proto!("kintai.v1");
        pub const KINTAI_SERVICE_FILE_DESCRIPTOR_SET: &[u8] =
            tonic::include_file_descriptor_set!("kintai_service_descriptor");
    }
}

pub mod notification {
    pub mod v1 {
        tonic::include_proto!("notification.v1");
        pub const NOTIFICATION_SERVICE_FILE_DESCRIPTOR_SET: &[u8] =
            tonic::include_file_descriptor_set!("notification_service_descriptor");
    }

    pub mod v2 {
        tonic::include_proto!("notification.v2");
        pub const NOTIFICATION_SERVICE_FILE_DESCRIPTOR_SET: &[u8] =
            tonic::include_file_descriptor_set!("v2_notification_service_descriptor");
    }
}

pub mod feed {
    pub mod v1 {
        tonic::include_proto!("feed.v1");
        pub const FEED_SERVICE_FILE_DESCRIPTOR_SET: &[u8] =
            tonic::include_file_descriptor_set!("v1_feed_service_descriptor");
    }
}

pub mod server_status {
    pub mod v1 {
        tonic::include_proto!("server_status.v1");
        pub const SERVER_STATUS_SERVICE_DESCRIPTOR_SET: &[u8] =
            tonic::include_file_descriptor_set!("server_status_service_descriptor");
    }
}
