function(ctx) {
  flow: {
    id: ctx.flow.id,
  },
  identity: {
    traits: {
      email: ctx.identity.traits.email,
    },
    // metadata_public: ctx.metadata_public,
  },
}
