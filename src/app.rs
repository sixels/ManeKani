use axum::Router;
use tokio::net::{TcpListener, ToSocketAddrs};

use crate::{routing::router::setup_router, state::AppState};

pub struct App {
    router: Router,
}

impl App {
    pub fn new(state: AppState) -> Self {
        let router = setup_router(state);
        Self { router }
    }

    pub async fn listen<S: ToSocketAddrs>(self, addr: S) -> std::io::Result<()> {
        let listener = TcpListener::bind(addr).await?;
        axum::serve(listener, self.router).await?;
        Ok(())
    }
}
