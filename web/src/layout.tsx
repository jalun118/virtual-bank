import { ReactNode } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import Navbar from "./components/Navbar";

const notIncludeNav = ["/login", "/register", "/"];
const notIncludeBack = ["/login", "/register", "/"];

export default function Layout({ children }: { children: ReactNode }) {
  const { pathname } = useLocation();
  const navigation = useNavigate();

  return (
    <div>
      {!notIncludeNav.includes(pathname) && <Navbar />}
      {!notIncludeBack.includes(pathname) && (
        <div className="m-auto mt-3 max-w-md">
          <div className="mb-5">
            <button
              onClick={() => navigation(-1)}
              className="flex items-center rounded-md px-2 py-1 text-lg font-semibold hover:bg-gray-200"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={1.5}
                stroke="currentColor"
                className="mr-2 size-6"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M10.5 19.5 3 12m0 0 7.5-7.5M3 12h18"
                />
              </svg>
              BACK
            </button>
          </div>
        </div>
      )}
      <main className="mt-3">{children}</main>
    </div>
  );
}
