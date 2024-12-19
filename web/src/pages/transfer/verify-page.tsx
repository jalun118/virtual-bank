import { useFormik } from "formik";
import { object as yupObject, string as yupString } from "yup";

interface iDefaultValue {
  destionation: string;
}

const defaultValue: iDefaultValue = {
  destionation: "",
};

const verifySchema = yupObject({
  destionation: yupString()
    .required("Required")
    .matches(/^[0-9]+$/, "Must be only number")
    .min(9, "Minimum 9 digit number"),
});

export default function VerifyPage({
  setPage,
}: {
  setPage: (page: number) => void;
}) {
  const { errors, handleChange, touched, handleSubmit, values } = useFormik({
    initialValues: defaultValue,
    onSubmit: (value) => {
      console.log(value);
      setPage(1);
      // action.setErrors({ destionation: "Mismatched accounts" });
    },
    validationSchema: verifySchema,
  });
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="destionation" className="mb-2 block font-semibold">
            Destionation Account
          </label>
          <input
            id="destionation"
            name="destionation"
            onChange={handleChange}
            value={values.destionation}
            autoComplete="off"
            autoCorrect="off"
            aria-autocomplete="none"
            spellCheck="false"
            className="block w-full rounded-md px-3 py-2 text-black"
            placeholder="7562844386140993"
          />
          {touched.destionation && errors.destionation && (
            <div className="font-semibold text-red-500">
              {errors.destionation}
            </div>
          )}
        </div>
        <div className="mt-5">
          <button
            type="submit"
            className="flex justify-between rounded-lg border bg-gradient-to-t from-cyan-500 to-cyan-300 px-3 py-2 font-semibold text-white shadow-md active:bg-gradient-to-b active:shadow-inner"
          >
            Verification
          </button>
        </div>
      </form>
    </div>
  );
}
