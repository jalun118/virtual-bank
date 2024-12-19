export default function DetailAccount({
  setPage,
}: {
  setPage: (page: number) => void;
}) {
  return (
    <div>
      <div className="flex items-center">
        <span className="mr-4 rounded-full bg-gray-200 p-1.5">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={1.5}
            stroke="currentColor"
            className="size-7"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z"
            />
          </svg>
        </span>
        <div>
          <span className="text-lg font-semibold">Ucup Surucup Bin Saha</span>
          <p className="text-sm font-semibold text-gray-600">NovaBank</p>
        </div>
      </div>
      <div className="mt-9 flex gap-3">
        <button
          onClick={() => setPage(2)}
          className="flex justify-between rounded-lg border bg-gradient-to-t from-cyan-500 to-cyan-300 px-3 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
        >
          Transfer
        </button>
        <button
          onClick={() => setPage(0)}
          className="flex justify-between rounded-lg border bg-gradient-to-t from-red-500 to-red-300 px-3 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
        >
          Cancel
        </button>
      </div>
    </div>
  );
}
