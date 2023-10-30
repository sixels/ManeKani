function(ctx) {
  user_id: ctx.identity.id,
  traits: {
    email: ctx.identity.traits.email,
    username: ctx.identity.traits.username,
  },
  created_at: ctx.identity.created_at,
}
