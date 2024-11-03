import { Link } from "react-router-dom";
import Brand from "./Brand";

export default function Navbar() {
  return (
    <header className="flex w-full flex-wrap bg-white py-3 md:container sm:flex-nowrap sm:justify-start">
      <nav className="mx-auto w-full max-w-[85rem] px-4 sm:flex sm:items-center sm:justify-between">
        <div className="flex items-center justify-between">
          <Link to="/">
            <Brand />
          </Link>
          <div className="inline-flex items-center gap-x-1 rounded-full bg-teal-100 px-3 py-1.5 font-medium text-teal-800 dark:bg-teal-500/10 dark:text-teal-500 sm:hidden">
            <span>Rp.</span>
            1000
          </div>
        </div>
        <div className="mt-5 flex flex-row items-center gap-5 sm:mt-0 sm:justify-end sm:ps-5">
          <Link
            className="font-medium text-gray-600 hover:text-gray-400 focus:text-gray-400 focus:outline-none dark:text-neutral-400 dark:hover:text-neutral-500 dark:focus:text-neutral-500"
            to=""
          >
            Home
          </Link>
          <Link
            className="font-medium text-gray-600 hover:text-gray-400 focus:text-gray-400 focus:outline-none dark:text-neutral-400 dark:hover:text-neutral-500 dark:focus:text-neutral-500"
            to="/account"
          >
            Account
          </Link>
          <Link
            className="font-medium text-gray-600 hover:text-gray-400 focus:text-gray-400 focus:outline-none dark:text-neutral-400 dark:hover:text-neutral-500 dark:focus:text-neutral-500"
            to="/top-up"
          >
            Top Up
          </Link>
          <Link
            className="font-medium text-gray-600 hover:text-gray-400 focus:text-gray-400 focus:outline-none dark:text-neutral-400 dark:hover:text-neutral-500 dark:focus:text-neutral-500"
            to="/login"
          >
            Login
          </Link>
          <Link
            className="font-medium text-gray-600 hover:text-gray-400 focus:text-gray-400 focus:outline-none dark:text-neutral-400 dark:hover:text-neutral-500 dark:focus:text-neutral-500"
            to="/register"
          >
            Register
          </Link>
          <div className="hidden items-center gap-x-1 rounded-full bg-teal-100 px-3 py-1.5 font-medium text-teal-800 dark:bg-teal-500/10 dark:text-teal-500 sm:inline-flex">
            <span>Rp.</span>
            1000
          </div>
        </div>
      </nav>
    </header>
  );
}
