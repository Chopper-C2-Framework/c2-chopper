import {Link} from "react-router-dom";

import  siteConfig  from "../config"
import { cn } from "@lib/utils"
import { Icons } from "@components/icons"
import {NavItem} from "../types.ts";

interface MainNavProps {
    items?: NavItem[]
}

export function MainNav({ items }: MainNavProps) {
    return (
        <div className="flex gap-6 md:gap-10">
            <Link to="/" className="hidden items-center space-x-2 md:flex">
                <Icons.logo className="h-12 w-12" />
                <span className="hidden font-bold sm:inline-block text-primary font-special">
          {siteConfig.name}
        </span>
            </Link>
            {items?.length ? (
                <nav className="hidden gap-6 md:flex">
                    {items?.map(
                        (item, index) =>
                            item.href && (
                                <Link
                                    key={index}
                                    to={item.href}
                                    className={cn(
                                        "flex items-center text-lg font-semibold text-muted-foreground sm:text-sm",
                                        item.disabled && "cursor-not-allowed opacity-80"
                                    )}
                                >
                                    {item.title}
                                </Link>
                            )
                    )}
                </nav>
            ) : null}
        </div>
    )
}

export default MainNav