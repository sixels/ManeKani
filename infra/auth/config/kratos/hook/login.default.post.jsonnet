function(ctx) {
  user_id: ctx.identity.id,
  email: ctx.identity.traits.email,
  created_at: ctx.identity.created_at,
}
