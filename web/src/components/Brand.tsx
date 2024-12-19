export default function Brand() {
  return (
    <div
      className="flex items-center text-2xl font-semibold"
      aria-label="Brand"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        strokeWidth={1.5}
        stroke="currentColor"
        className="mr-2 size-10 rounded-full bg-blue-700 p-1.5 text-white"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          d="M2.25 8.25h19.5M2.25 9h19.5m-16.5 5.25h6m-6 2.25h3m-3.75 3h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5Z"
        />
      </svg>
      <span className="italic text-blue-700">NovaBank</span>
    </div>
  );
}

export function BrandDark({
  fontSize,
  iconSize,
}: {
  fontSize?: string;
  iconSize?: string;
}) {
  return (
    <div
      className={`flex items-center ${fontSize ?? "text-2xl"} font-semibold`}
      aria-label="Brand"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        strokeWidth={1.5}
        stroke="currentColor"
        className={`mr-2 ${iconSize ?? "size-10"} rounded-full bg-white p-1 text-blue-700`}
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          d="M2.25 8.25h19.5M2.25 9h19.5m-16.5 5.25h6m-6 2.25h3m-3.75 3h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5Z"
        />
      </svg>
      <span className="italic text-white">NovaBank</span>
    </div>
  );
}
