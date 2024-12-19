import { Link } from "react-router-dom";
import { BrandDark } from "../../components/Brand";
import { CurrencyRupiahFormat } from "../../utils/currency";

function CreateSpaceText(text: string, space_every: number): string {
  let resultText = "";
  let indexSpace = 0;
  for (let index = 0; index < text.length; index++) {
    const char = text[index];
    if (index === indexSpace * space_every) {
      resultText += " " + char;
      indexSpace += 1;
    } else {
      resultText += char;
    }
  }
  return resultText;
}

export default function MyAccount() {
  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-2xl font-semibold">Account</h1>
        <div className="mt-4">
          <div className="w-full rounded-xl bg-gradient-to-tr from-blue-500 via-indigo-500 via-50% to-cyan-500 px-4 py-4 text-white">
            <div className="mb-4">
              <BrandDark fontSize="text-lg" iconSize="size-7" />
            </div>
            <div className="px-2 font-mono">
              <h2 className="select-none text-xl tabular-nums">
                {CurrencyRupiahFormat(200000)}
              </h2>
              <p className="mt-20 select-none text-right text-lg tabular-nums">
                {CreateSpaceText("5440290249006276", 4)}
              </p>
            </div>
          </div>
          <div className="mt-4 flex gap-3">
            <Link
              to="/top-up"
              className="flex flex-col items-center rounded-xl border border-gray-400 px-1.5 py-2 hover:bg-gray-100"
            >
              <span className="rounded-full bg-green-500 p-1.5 text-white shadow-md">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={1.5}
                  stroke="currentColor"
                  className="size-6"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M12 4.5v15m7.5-7.5h-15"
                  />
                </svg>
              </span>
              <span className="mt-1.5 font-semibold">Top Up</span>
            </Link>
            <button className="flex flex-col items-center rounded-xl border border-gray-400 px-1.5 py-2 hover:bg-gray-100">
              <span className="rounded-full bg-cyan-500 p-1.5 text-white shadow-md">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  strokeWidth={1.5}
                  stroke="currentColor"
                  className="size-6"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    d="M7.5 21 3 16.5m0 0L7.5 12M3 16.5h13.5m0-13.5L21 7.5m0 0L16.5 12M21 7.5H7.5"
                  />
                </svg>
              </span>
              <span className="mt-1.5 font-semibold">Transfer</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
