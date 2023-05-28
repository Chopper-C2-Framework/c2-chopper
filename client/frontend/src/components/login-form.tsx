import * as React from "react";
import * as z from "zod";

import { Button } from "@components/ui/button";
import { Input } from "@components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { cn } from "@lib/utils";
import { AxiosResponse } from "axios";
import { useForm } from "react-hook-form";
import { UseMutationResult } from "react-query";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "./ui/form";

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {
  useMutation: () => UseMutationResult<
    AxiosResponse<any, any>,
    unknown,
    any,
    unknown
  >;
  formSchema: any;
  defaultValues: any;
}

export function AuthForm({
  className,
  defaultValues,
  useMutation,
  formSchema,
  ...props
}: UserAuthFormProps) {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues,
  });

  const { mutate, error, data, isLoading } = useMutation();

  async function onSubmit(data:any) {
    mutate({ username:data.username, password:data.password });
  }

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input placeholder="h3x00rr_cr4t" {...field} />
                </FormControl>
                <FormDescription>
                  This is your public display name.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

  <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input placeholder="password" type="password" {...field} />
                </FormControl>
                
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit">Submit</Button>
        </form>
      </Form>
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t" />
        </div>
        <div className="relative flex justify-center text-xs uppercase"></div>
      </div>
    </div>
  );
}
