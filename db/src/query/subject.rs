mod create_subject;
mod get_subject;
mod update_subject;

pub use self::{
    create_subject::{create_subject, CreateSubject},
    get_subject::{get_deck_subjects, get_subject, GetDeckSubjectsFilter},
    // update_subject::{update_subject, UpdateSubject},
};

// TODO: update, bulk create, delete, etc.
