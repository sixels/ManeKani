function(ctx) {
  created_at: ctx.identity.state_changed_at,
  user_id: ctx.identity.id,
  email: ctx.identity.traits.email,
}
