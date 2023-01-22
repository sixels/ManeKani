import AuthRoute from "@/lib/auth/wrappers/AuthRoute";

function Dashboard() {
  return (
    <>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
      Abacate
    </>
  );
}

export default function Wrapper() {
  return (
    <AuthRoute>
      <Dashboard />
    </AuthRoute>
  );
}
