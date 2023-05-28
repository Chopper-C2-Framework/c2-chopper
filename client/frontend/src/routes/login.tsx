import { AuthForm } from "@components/login-form";
import { loginSchema } from "@hooks/mutations/auth/login";
import { useRegisterMutation } from "@hooks/mutations/auth/register";
export default function Login() {
  return (
    <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Login to your account
        </h1>
        <p className="text-sm text-muted-foreground">
          Enter your username and password below to login
        </p>
      </div>
      <AuthForm useMutation={useRegisterMutation} defaultValues={{username:"h3x004_cr4t",password:"xxxxxx"}} formSchema={loginSchema}/>
    </div>
  );
}
