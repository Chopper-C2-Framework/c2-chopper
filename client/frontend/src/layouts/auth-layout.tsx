import React from 'react'
import {Link, Outlet, useLocation} from "react-router-dom"
import {cn} from "@lib/utils"
import {buttonVariants} from "@components/ui/button"
import {Footer} from "@components/footer.tsx";
import { Icons } from "@components/icons"
import WallpaperAuth2 from "../assets/wallpaper-auth-2.jpg"


import WallpaperAuth from "../assets/wallpaper-auth.jpg"

interface AuthLayoutProps {

}

export const AuthLayout: React.FC<AuthLayoutProps> = ({}) => {
    const { pathname}=useLocation()
    console.log(pathname)
    return (
  <div className="w-full h-full">
            <div className="md:hidden ">
                <img
                    src="/examples/authentication-light.png"
                    width={1280}
                    height={843}
                    alt="Authentication"
                    className="block dark:hidden"
                />
                <img
                    src="/examples/authentication-dark.png"
                    width={1280}
                    height={843}
                    alt="Authentication"
                    className="hidden dark:block"
                />
            </div>
            <div className="container relative hidden h-[900px] flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2  px-20 py-5">
                <Link to={pathname==="/login"?"/register":"/login"}
                    className={cn(
                        buttonVariants({ variant: "ghost", size: "sm" }),
                        "absolute right-4 top-4 md:right-8 md:top-8"
                    )}
                >
                    {pathname === "/login" ? "Register" : "Login"}
                </Link>
                <div className="relative hidden h-[80vh] flex-col bg-secondary-foreground p-10 text-white dark:border-r lg:flex">
                    <div
                        className="absolute inset-0 bg-cover  bg-center bg-no-repeat " 
                        style={{
                            backgroundImage:
                                `url(${pathname==="/login"?WallpaperAuth:WallpaperAuth2})`,
                        }}
                    />
                    <Link to="/">
                    <div className="relative z-20 flex items-center text-lg font-medium">
                        <Icons.logo className="mr-2 h-10 w-10" /> <span className="text-primary font-black font-special">C2-Chopper</span>
                    </div>
</Link>
                    <div className="relative z-20 mt-auto">
                        <blockquote className="space-y-2">
                            <p className="text-lg text-white font-bold  text-3xl italic font-special bg-black w-fit text-center">
                             Evolves with your team, Hack your way to the top, and break into the system faster than ever before.
                            </p>
                        </blockquote>
                    </div>
                </div>
                <div className="lg:p-8">
                  
                        <Outlet/>
                </div>
            </div>
            <div className="mt-20">
            <Footer/>
            </div>
        </div>
        );
}


export default AuthLayout