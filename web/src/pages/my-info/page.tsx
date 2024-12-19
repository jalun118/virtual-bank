import { Link } from "react-router-dom";

export default function MyInfo() {
  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-xl font-semibold">My Information</h1>
        <div className="mt-3">
          <div>
            <span className="font-semibold text-gray-700">Email</span>
            <p className="mt-1 border-b border-b-gray-400 pb-2 text-base">
              kuda@gmail.com
            </p>
          </div>
          <div className="mt-2">
            <span className="font-semibold text-gray-700">Username</span>
            <p className="mt-1 border-b border-b-gray-400 pb-2 text-base">
              kuda_jemping434
            </p>
          </div>
          <div className="mt-2">
            <span className="font-semibold text-gray-700">Full Name</span>
            <p className="mt-1 border-b border-b-gray-400 pb-2 text-base">
              kuda jemping
            </p>
          </div>
          <div className="mt-2">
            <span className="font-semibold text-gray-700">Birth Date</span>
            <p className="mt-1 border-b border-b-gray-400 pb-2 text-base tabular-nums">
              01/10/1984
            </p>
          </div>
        </div>
        <div className="mt-4 flex gap-x-2">
          <Link
            to="/my-info/edit"
            className="rounded-md bg-cyan-600 px-3 py-2 font-semibold text-white"
          >
            Change Username
          </Link>
          <Link
            to="/my-info/change-password"
            className="rounded-md bg-emerald-600 px-3 py-2 font-semibold text-white"
          >
            Change Password
          </Link>
        </div>
      </div>
    </div>
  );
}
