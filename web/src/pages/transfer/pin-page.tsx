import {
  ChangeEvent,
  KeyboardEvent,
  useCallback,
  useEffect,
  useRef,
  useState,
} from "react";

const initialValuePin = new Array(6).fill("");

export default function PinPage({
  setPage,
}: {
  setPage: (page: number) => void;
}) {
  const [pin, setPin] = useState(initialValuePin);
  const inputRefs = useRef<(HTMLInputElement | null)[]>([]);

  function handleChange(e: ChangeEvent<HTMLInputElement>, index: number) {
    const value = e.currentTarget.value;
    setPin((prev) => {
      const newPin = [...prev];

      newPin[index] = value.substring(value.length - 1);
      return newPin;
    });

    if (
      value &&
      index < inputRefs.current.length - 1 &&
      inputRefs.current[index + 1]
    ) {
      inputRefs.current[index + 1]?.focus();
    }
  }

  function handleClick(index: number) {
    inputRefs.current[index]?.setSelectionRange(1, 1);

    if (index > 0 && !pin[index - 1]) {
      inputRefs.current[pin.indexOf("")]?.focus();
    }
  }

  function handleKeyDown(e: KeyboardEvent, index: number) {
    if (
      e.key === "Backspace" &&
      !pin[index] &&
      index > 0 &&
      inputRefs.current[index - 1]
    ) {
      inputRefs.current[index - 1]?.focus();
    }
  }

  const handleSubmit = useCallback(() => {
    console.log("dasda");

    setPage(4);
  }, [setPage]);

  useEffect(() => {
    const pinResult = pin.join("").trim();
    if (pinResult.length === pin.length) {
      handleSubmit();
    }
  }, [pin, handleSubmit]);

  return (
    <div>
      <div className="flex gap-4">
        {pin.map((data, i) => (
          <input
            key={i}
            autoFocus={i === 0}
            autoComplete="off"
            type="password"
            autoCorrect="off"
            aria-autocomplete="none"
            spellCheck="false"
            ref={(input) => (inputRefs.current[i] = input)}
            maxLength={1}
            value={data}
            onKeyDown={(e) => handleKeyDown(e, i)}
            onChange={(e) => handleChange(e, i)}
            onClick={() => handleClick(i)}
            className="w-full rounded-md px-0 text-center text-lg text-black"
          />
        ))}
      </div>
      <div className="mt-3 font-semibold text-red-500">
        Password Not Matcher
      </div>
      <div className="mt-7">
        <button
          onClick={() => handleSubmit()}
          className="flex justify-between rounded-lg border bg-gradient-to-t from-cyan-500 to-cyan-300 px-3 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
        >
          Transfer
        </button>
      </div>
    </div>
  );
}
