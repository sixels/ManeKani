// import Image from "next/image";

// import Logo from "@/assets/logo-aside-dark.svg";

export default function Footer() {
  return (
    <div className="bg-neutral-800 text-neutral-300  w-full">
      <div className="flex flex-col  py-4 pt-8 gap-4 justify-center mx-auto max-w-screen-2xl items-center">
        <div className="py-6">
          <img
            src="/assets/icon-196.png"
            alt="ManeKani Logo"
            className="w-[150px]"
          />
        </div>
        <div className="flex flex-row gap-6">
          <a className="hover:underline" href="/">
            Home
          </a>
          <a className="hover:underline" href="/dashboard">
            Dashboard
          </a>
          <a className="hover:underline" href="/about">
            About
          </a>
          <a className="hover:underline" href="/contact">
            Contact
          </a>
        </div>
      </div>

      <div className="border-t mx-auto max-w-screen-2xl  border-neutral-700">
        <div className="flex flex-col md:flex-col gap-4 py-4 justify-between items-center">
          <p>© 2023 ManeKani. All rights reserved</p>
          <div className="flex flex-row gap-6" />
        </div>
      </div>
    </div>
  );
}
