import { useUser } from "@/lib/hooks/user";
import Dashboard from "./dashboard";

export default function Index() {
  const { user, isError, isLoading } = useUser();

  if (isLoading) {
    return <></>;
  }
  if (isError || !user) {
    return <>Index Page</>;
  }

  return <Dashboard />;
}
