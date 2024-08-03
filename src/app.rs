use std::{fmt::Debug, net::ToSocketAddrs};

use actix_web::{
    guard,
    middleware::{NormalizePath, TrailingSlash},
    web, App, HttpServer,
};

use crate::{
    routing::router::{graphiql, ManeKaniGQLApi, SettingsRouter},
    state::AppState,
};

pub struct ManeKani {
    state: AppState,
}

impl ManeKani {
    pub fn new(state: AppState) -> Self {
        Self { state }
    }

    pub async fn listen<A: ToSocketAddrs + Debug>(self, addr: A) -> std::io::Result<()> {
        let AppState { db, auth_manager } = self.state.clone();

        HttpServer::new(move || {
            App::new()
                .wrap(NormalizePath::trim())
                .app_data(web::Data::from(db.clone()))
                .app_data(web::Data::from(auth_manager.clone()))
                .service(actix_files::Files::new("/assets", "./assets"))
                .service(SettingsRouter.scope())
                .service(web::resource("/graphiql").guard(guard::Get()).to(graphiql))
                .setup_graphql(db.clone())
        })
        .bind(&addr)
        .inspect(|_| {
            println!("Server started at http://{:?}", addr);
        })?
        .run()
        .await?;

        Ok(())
    }
}
