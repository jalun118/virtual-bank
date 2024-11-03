import { Link } from "react-router-dom";
import Brand from "../../components/Brand";

export default function LoginPage() {
  return (
    <div className="mt-20 flex w-full justify-center">
      <div className="min-h-[28%] w-[90%] rounded-2xl border bg-white p-8 shadow-md md:w-[50%] lg:w-[40%] xl:w-[28%]">
        <div className="flex justify-center">
          <Brand />
        </div>
        <form>
          <h3 className="text-center text-2xl">Login</h3>
          <div className="mt-3 max-w-sm">
            <label
              htmlFor="email"
              className="mb-2 block text-sm font-medium dark:text-white"
            >
              Email
            </label>
            <input
              type="email"
              autoComplete="off"
              autoCorrect="off"
              aria-autocomplete="none"
              id="email"
              className="block w-full rounded-lg border-gray-200 px-3 py-3 text-sm focus:border-blue-500 focus:ring-blue-500 disabled:pointer-events-none disabled:opacity-50 dark:border-neutral-700 dark:bg-neutral-900 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600"
              placeholder="you@site.com"
            />
          </div>

          <div className="mt-3 max-w-sm">
            <label
              htmlFor="password"
              className="mb-2 block text-sm font-medium dark:text-white"
            >
              Password
            </label>
            <input
              type="text"
              autoComplete="off"
              autoCorrect="off"
              aria-autocomplete="none"
              id="password"
              className="block w-full rounded-lg border-gray-200 px-3 py-3 text-sm focus:border-blue-500 focus:ring-blue-500 disabled:pointer-events-none disabled:opacity-50 dark:border-neutral-700 dark:bg-neutral-900 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600"
              placeholder="*********"
            />
          </div>
          <p className="mt-2 text-base">
            Don't have an account yet?{" "}
            <Link to="/register" replace className="text-blue-500">
              register
            </Link>
          </p>

          <div className="mt-3 flex justify-end">
            <button
              type="button"
              className="inline-flex items-center gap-x-2 rounded-lg border border-transparent bg-blue-600 px-4 py-2.5 font-medium text-white hover:bg-blue-700 focus:bg-blue-700 focus:outline-none disabled:pointer-events-none disabled:opacity-50"
            >
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
