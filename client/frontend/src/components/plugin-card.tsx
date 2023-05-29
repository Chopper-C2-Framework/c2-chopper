import { Github, Info, ToyBrick } from "lucide-react";
import React from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { usePluginsDetails } from "@hooks/queries/plugins/plugin-detail";
import { Badge } from "./ui/badge";

interface PluginCardProps {
  plugin: string;
}

export const PluginCard: React.FC<PluginCardProps> = ({ plugin }) => {
  return (
    <Card className="col-span-3">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium font-black text-lg text-primary">
          {plugin}
        </CardTitle>
        <ToyBrick className="w-6 h-6 text-primary" />
      </CardHeader>
    </Card>
  );
};

export const LoadedPluginCard: React.FC<PluginCardProps> = ({ plugin }) => {
  const { data, error } = usePluginsDetails(plugin);
  return (
    <Card className="col-span-3">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium font-black text-lg text-primary">
          {data?.info.Name}
        </CardTitle>
        <div className="space-x-2 flex">
          <ToyBrick className="w-6 h-6 text-primary" />
          <a target="_blank" href={data?.Metadata.source_link} className="cursor-pointer hover:opacity-80">
            <Github className="w-6 h-6 text-secondary" />
          </a>
        </div>
      </CardHeader>
      <CardDescription className="px-8 text-left ">
        {data?.Metadata.description}
      </CardDescription>
      <div className="h-4"></div>

      <CardContent>
        <InfoLabel label="Author" info={data?.Metadata.author} />
        <InfoLabel label="Version" info={data?.Metadata.version} />
        <InfoLabel label="Release Date" info={data?.Metadata.release_date} />
        <InfoLabel
          label="Type"
          info={mapTypeNumberToName(data?.Metadata.type)}
        />
      </CardContent>
      <CardFooter className="px-8">
        <PluginTags tags={data?.Metadata.tags || []} />
      </CardFooter>
    </Card>
  );
};

interface InfoLabelProps {
  label?: string;
  info?: string;
}

const InfoLabel: React.FC<InfoLabelProps> = ({ info, label }) => (
  <div className="flex items-center space-x-8 px-2">
    <p className="font-bold text-primary/90 text-md">{label}</p>
    <p>{info}</p>
  </div>
);

interface PluginTagsProps {
  tags: string[];
}

const PluginTags: React.FC<PluginTagsProps> = ({ tags }) => {
  return (
    <div className="flex items-center space-x-5 wrap">
      {tags.map((t) => (
        <Badge variant="default">{t}</Badge>
      ))}
    </div>
  );
};

const mapTypeNumberToName = (type?: number) => {
  switch (type) {
    case 0:
      return "Info Retriever";
    case 1:
      return "Session Opener";
  }
};
