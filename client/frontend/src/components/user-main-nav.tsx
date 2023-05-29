import { Link } from "react-router-dom";
import { cn } from "@lib/utils";

export function UserMainNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <nav
      className={cn("flex items-center space-x-4 lg:space-x-6", className)}
      {...props}
    >
      <Link
        to="/app/dashboard"
        className="text-sm font-medium transition-colors hover:text-primary"
      >
        Overview
      </Link>
      <Link
        to="/app/plugins"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Plugins
      </Link>
      <Link
        to="/app/findings"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Findings
      </Link>
      <Link
        to="/app/tasks"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Tasks
      </Link>
      <Link
        to="/app/report"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Report
      </Link>
    </nav>
  );
}
