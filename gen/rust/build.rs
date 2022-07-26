const PROTO_ROOT_DIR: &str = "../../proto";

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let out_dir = std::path::PathBuf::from(std::env::var("OUT_DIR")?);
    tonic_build::configure()
        .protoc_arg("--experimental_allow_proto3_optional")
        .file_descriptor_set_path(out_dir.join("kintai_service_descriptor.bin"))
        .compile(
            &[format!("{PROTO_ROOT_DIR}/kintai/v1/kintai_service.proto")],
            &[PROTO_ROOT_DIR],
        )?;
    tonic_build::configure()
        .protoc_arg("--experimental_allow_proto3_optional")
        .file_descriptor_set_path(out_dir.join("notification_service_descriptor.bin"))
        .compile(
            &[format!(
                "{PROTO_ROOT_DIR}/notification/v1/notification_service.proto"
            )],
            &[PROTO_ROOT_DIR],
        )?;
    tonic_build::configure()
        .protoc_arg("--experimental_allow_proto3_optional")
        .file_descriptor_set_path(out_dir.join("release_service_descriptor.bin"))
        .compile(
            &[format!(
                "{PROTO_ROOT_DIR}/automation/v1/release_service.proto"
            )],
            &[PROTO_ROOT_DIR],
        )?;
    tonic_build::configure()
        .protoc_arg("--experimental_allow_proto3_optional")
        .file_descriptor_set_path(out_dir.join("server_status_service_descriptor.bin"))
        .compile(
            &[format!(
                "{PROTO_ROOT_DIR}/server_status/v1/server_status_service.proto"
            )],
            &[PROTO_ROOT_DIR],
        )?;
    Ok(())
}
