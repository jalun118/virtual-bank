import { Link } from "react-router-dom";
import Brand from "../../components/Brand";

export default function HomePage() {
  return (
    <div className="mt-20 flex w-full justify-center">
      <div className="min-h-[28%] w-[90%] rounded-2xl border bg-white p-8 shadow-md md:w-[50%] lg:w-[40%] xl:w-[28%]">
        <div className="flex justify-center">
          <Brand />
        </div>
        <div className="mt-3">
          <div className="m-auto mt-7 flex max-w-72 flex-col gap-4">
            <Link
              to="/my-info"
              className="flex justify-center rounded-lg border border-gray-400 p-3 text-lg hover:bg-gray-100 active:bg-gray-200"
            >
              <div className="flex items-center gap-2">My Info</div>
            </Link>
            <Link
              to="/my-account"
              className="flex justify-center rounded-lg border border-gray-400 p-3 text-lg hover:bg-gray-100 active:bg-gray-200"
            >
              <div className="flex items-center gap-2">My Account</div>
            </Link>
            <Link
              to="/transaction"
              className="flex justify-center rounded-lg border border-gray-400 p-3 text-lg hover:bg-gray-100 active:bg-gray-200"
            >
              <div className="flex items-center gap-2">Transaction</div>
            </Link>
            <Link
              to="/transfer"
              className="flex justify-center rounded-lg border border-gray-400 p-3 text-lg hover:bg-gray-100 active:bg-gray-200"
            >
              <div className="flex items-center gap-2">Transfer</div>
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
