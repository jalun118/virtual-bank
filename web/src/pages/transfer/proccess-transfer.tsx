import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

export default function ProccessTransfer() {
  const navigate = useNavigate();

  useEffect(() => {
    const st = setTimeout(() => {
      navigate("/transaction/73531d72-169d-4ce8-b7fb-625161fcbd82", {
        replace: true,
      });
    }, 3000);

    return () => clearTimeout(st);
  }, [navigate]);

  return (
    <div>
      <div className="flex flex-col items-center overflow-hidden pt-20">
        <div className="animate-[backInUp_2s_cubic-bezier(0.72,-0.71,0.34,1.5)] rounded-full bg-gradient-to-tr from-cyan-400 to-emerald-400 p-3 text-white">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            strokeWidth={3}
            stroke="currentColor"
            className="size-16"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              d="m4.5 12.75 6 6 9-13.5"
            />
          </svg>
        </div>
        <p className="mt-5 animate-[lightSpeedInLeft_2s_cubic-bezier(0.72,-0.71,0.34,1.5)] text-2xl font-semibold">
          Success Transfer
        </p>
      </div>
    </div>
  );
}
