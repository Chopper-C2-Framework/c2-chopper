"use client";

import { Check, ChevronsUpDown, PlusCircle } from "lucide-react";
import * as React from "react";

import { Avatar, AvatarFallback, AvatarImage } from "@components/ui/avatar";
import { Button } from "@components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
} from "@components/ui/command";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@components/ui/dialog";
import { Input } from "@components/ui/input";
import { Label } from "@components/ui/label";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@components/ui/popover";
import { useCreateTeam } from "@hooks/mutations/team/create-team";
import { cn } from "@lib/utils";
import { useGetTeams } from "@hooks/queries/team/getTeams";

const groups = [
  {
    label: "Teams",
    teams: [
      {
        label: "Acme Inc.",
        value: "acme-inc",
      },
      {
        label: "Monsters Inc.",
        value: "monsters",
      },
    ],
  },
];

type PopoverTriggerProps = React.ComponentPropsWithoutRef<
  typeof PopoverTrigger
>;

interface TeamSwitcherProps extends PopoverTriggerProps {}

export default function TeamSwitcher({ className }: TeamSwitcherProps) {
  const { data: allTeamsData, isLoading: alLTeamsLoading } = useGetTeams();

  const [selectedTeam, setTeam] = React.useState<any>({
    name: "loading",
    members: [],
  });

  const localGroups = allTeamsData
    ? [
        {
          label: "Teams",
          teams: allTeamsData
            ? allTeamsData.teams.map((t) => ({
                label: t,
                value: t,
              }))
            : [],
        },
      ]
    : groups;

  const [open, setOpen] = React.useState(false);
  const [showNewTeamDialog, setShowNewTeamDialog] = React.useState(false);

  console.log(allTeamsData);

  return (
    <Dialog open={showNewTeamDialog} onOpenChange={setShowNewTeamDialog}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            size="sm"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a team"
            className={cn("w-[200px] justify-between", className)}
          >
            <ChevronsUpDown className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command>
            <CommandList>
              <CommandInput placeholder="Search team..." />
              <CommandEmpty>No team found.</CommandEmpty>
              {localGroups.map((group) => (
                <CommandGroup key={group.label} heading={group.label}>
                  {allTeamsData?.teams.map((team, index) => (
                    <CommandItem
                      key={team.id}
                      onSelect={() => {
                        setTeam(team);
                        setOpen(false);
                        console.log("here");
                      }}
                      className="text-sm"
                    >
                      {team.name}
                      <Check className={cn("ml-auto h-4 w-4", "opacity-100")} />
                    </CommandItem>
                  ))}
                </CommandGroup>
              ))}
            </CommandList>
            <CommandSeparator />
            <CommandList>
              <CommandGroup>
                <DialogTrigger asChild>
                  <CommandItem
                    onSelect={() => {
                      setOpen(false);
                      setShowNewTeamDialog(true);
                    }}
                  >
                    <PlusCircle className="mr-2 h-5 w-5" />
                    Create Team
                  </CommandItem>
                </DialogTrigger>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
      <CreateTeamDialog setShowNewTeamDialog={setShowNewTeamDialog} />
    </Dialog>
  );
}

interface CreateTeamDialoagProps {
  setShowNewTeamDialog: (value: boolean) => void;
}
export const CreateTeamDialog: React.FC<CreateTeamDialoagProps> = ({
  setShowNewTeamDialog,
}) => {
  const [teamName, setTeamName] = React.useState("");
  const { mutateAsync, data, isLoading } = useCreateTeam();
  const onSubmit = async () => {
    await mutateAsync({ name: teamName });
    setShowNewTeamDialog(false);
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Create team</DialogTitle>
        <DialogDescription>
          Add a new team to manage your operations and projects.
        </DialogDescription>
      </DialogHeader>
      <div>
        <div className="space-y-4 py-2 pb-4">
          <div className="space-y-2">
            <Label htmlFor="name">Team name</Label>
            <Input
              id="name"
              placeholder="Acme Inc."
              onChange={({ target }) => setTeamName(target.value)}
            />
          </div>
        </div>
      </div>
      <DialogFooter>
        <Button variant="outline" onClick={() => setShowNewTeamDialog(false)}>
          Cancel
        </Button>
        <Button type="submit" onClick={onSubmit}>
          Continue
        </Button>
      </DialogFooter>
    </DialogContent>
  );
};
