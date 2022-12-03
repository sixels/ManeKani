import Header from '@/ui/Header';

import './style.css';

type LayoutProps = { children: React.ReactNode };
export default function Layout({ children }: LayoutProps) {
  return (
    <>
      <Header />
      {children}
    </>
  );
}
