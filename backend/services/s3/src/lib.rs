#![warn(clippy::pedantic)]

pub mod domain;
pub mod model;
pub mod repository;

pub use repository::S3Repo;
