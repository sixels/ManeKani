function(ctx) {
  flow: {
    id: ctx.flow.id,
  },
  identity: {
    traits: {
      email: ctx.identity.traits.email,
    },
  },
}
