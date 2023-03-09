import Header from "@/components/Header";
import Link from "next/link";
import React from "react";

const NotFound = () => {
  return (
    <>
      <Header />
      <div className="flex flex-col gap-10 justify-center items-center h-[85vh]">
        <h1 className="text-4xl">404 Not Found !</h1>
        <Link href="/">
            Got To Home Page
        </Link>
      </div>
    </>
  );
};

export default NotFound;