import { useFormik } from "formik";
import { Link } from "react-router-dom";
import * as Yup from "yup";
import Brand from "../../components/Brand";

const loginSchema = Yup.object().shape({
  email: Yup.string().email("Invalid email").required("Required"),
  password: Yup.string().required("Required"),
});

interface iInitValues {
  email: string;
  password: string;
}

const initialValue: iInitValues = {
  email: "",
  password: "",
};

export default function LoginPage() {
  const { errors, values, touched, handleChange, handleSubmit } = useFormik({
    initialValues: initialValue,
    onSubmit: (values) => {
      console.log(values);
      window.location.replace("/");
    },
    validationSchema: loginSchema,
  });

  return (
    <div className="mt-20 flex w-full justify-center">
      <div className="min-h-[28%] w-[90%] rounded-2xl border bg-white p-8 shadow-md md:w-[50%] lg:w-[40%] xl:w-[28%]">
        <div className="flex justify-center">
          <Brand />
        </div>
        <form onSubmit={handleSubmit}>
          <h3 className="mt-3 text-center text-2xl font-semibold">Login</h3>
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
              name="email"
              value={values.email}
              onChange={handleChange}
              className="block w-full rounded-lg border-gray-200 px-3 py-3 text-sm focus:border-blue-500 focus:ring-blue-500 disabled:pointer-events-none disabled:opacity-50 dark:border-neutral-700 dark:bg-neutral-900 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600"
              placeholder="you@site.com"
            />
            {touched.email && errors.email && (
              <div className="mt-1 text-sm font-semibold text-red-500">
                {errors.email}
              </div>
            )}
          </div>

          <div className="mt-3 max-w-sm">
            <label
              htmlFor="password"
              className="mb-2 block text-sm font-medium dark:text-white"
            >
              Password
            </label>
            <input
              type="password"
              autoComplete="off"
              autoCorrect="off"
              aria-autocomplete="none"
              id="password"
              name="password"
              value={values.password}
              onChange={handleChange}
              className="block w-full rounded-lg border-gray-200 px-3 py-3 text-sm focus:border-blue-500 focus:ring-blue-500 disabled:pointer-events-none disabled:opacity-50 dark:border-neutral-700 dark:bg-neutral-900 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600"
              placeholder="*********"
            />
            {touched.password && errors.password && (
              <div className="mt-1 text-sm font-semibold text-red-500">
                {errors.password}
              </div>
            )}
          </div>
          <p className="mt-2 text-base">
            Don't have an account yet?{" "}
            <Link to="/register" replace className="text-blue-500">
              register
            </Link>
          </p>

          <div className="mt-3 flex justify-end">
            <button
              type="submit"
              className="items-center rounded-lg border border-transparent bg-blue-600 px-4 py-2 font-medium text-white hover:bg-blue-700 focus:bg-blue-700 focus:outline-none disabled:pointer-events-none disabled:opacity-50"
            >
              Login
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
