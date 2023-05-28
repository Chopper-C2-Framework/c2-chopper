import { Outlet } from "react-router-dom";
import { Link } from "react-router-dom";
import { Footer } from "@components/footer.tsx";
import { cn } from "@lib/utils.ts";
import { buttonVariants } from "@components/ui/button.tsx";
import config from "../config.ts";
import MainNav from "@components/main-nav.tsx";
import { useMeQuery } from "@hooks/queries/user/me";

export function NavbarLayout() {
  const { isError, isLoading } = useMeQuery();
  return (
    <div className="flex min-h-screen flex-col">
      <header className="container z-40 bg-background">
        <div className="flex h-20 items-center justify-between py-6">
          <MainNav items={config.mainNav} />

          <nav>
            {isLoading && !isError && (
              <Link
                to="/login"
                className={cn(
                  buttonVariants({ variant: "secondary", size: "sm" }),
                  "px-4"
                )}
              >
                Login
              </Link>
            )}
          </nav>
        </div>
      </header>
      <main className="flex-1">
        <Outlet />
      </main>
      <Footer />
    </div>
  );
}

export default NavbarLayout;
