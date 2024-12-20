import { useFormik } from "formik";
import { object as yupObject, string as yupString } from "yup";

interface defaultValue {
  email: string;
}

const initValue: defaultValue = {
  email: "",
};

const lostPasswordSchema = yupObject({
  email: yupString().required("Required").email("Invalid email"),
});

export default function ForgotPasswordPage() {
  const { errors, touched, handleChange, handleSubmit, values } = useFormik({
    initialValues: initValue,
    onSubmit: (value, action) => {
      console.log(value);
      action.resetForm();
    },
    validationSchema: lostPasswordSchema,
  });

  return (
    <div className="container">
      <div className="m-auto max-w-md">
        <h1 className="text-2xl font-semibold">Forgot Password</h1>
        <div className="mt-4">
          <form onSubmit={handleSubmit}>
            <p>
              You will be provided with a link to reset your password via a
              valid email.
            </p>
            <div className="mt-4">
              <input
                type="email"
                autoComplete="off"
                autoCorrect="off"
                aria-autocomplete="none"
                spellCheck="false"
                value={values.email}
                onChange={handleChange}
                name="email"
                className="block w-full rounded-md px-3 py-2 pe-9 text-black"
                placeholder="your@email.com"
              />
              {errors.email && touched.email && (
                <div className="font-semibold text-red-500">{errors.email}</div>
              )}
            </div>
            <div className="mt-4">
              <button
                type="submit"
                className="rounded-md bg-blue-600 px-4 py-2 font-semibold text-white active:bg-blue-700"
              >
                Send
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
}
