import { Link } from "react-router-dom";
import Brand from "./Brand";

export default function Navbar() {
  return (
    <header className="m-auto flex w-full max-w-md flex-wrap bg-white py-3 sm:flex-nowrap sm:justify-start">
      <nav className="mx-auto w-full max-w-[85rem] px-4 sm:flex sm:items-center sm:justify-between">
        <div className="flex items-center justify-between">
          <Link to="/">
            <Brand />
          </Link>
        </div>
      </nav>
    </header>
  );
}
