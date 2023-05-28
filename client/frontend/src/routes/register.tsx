import { AuthForm } from "@components/login-form";
import { useRegisterMutation, registerSchema } from "@hooks/mutations/auth/register";
export default function Register() {
  return (
    <div className="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
      <div className="flex flex-col space-y-2 text-center">
        <h1 className="text-2xl font-semibold tracking-tight">
          Create an account
        </h1>
        <p className="text-sm text-muted-foreground">
          Enter your username below to create your account
        </p>
      </div>
      <AuthForm  defaultValues={{username:"h3x004_cr4t",password:"xxxxxx"}} formSchema={registerSchema} useMutation={useRegisterMutation}/>
    </div>
  );
}

