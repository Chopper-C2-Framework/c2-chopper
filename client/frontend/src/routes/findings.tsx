import Heading from "@components/heading";
import { useAllCreds } from "@hooks/queries/tracking/creds";
import { useAllHosts } from "@hooks/queries/tracking/hosts";
import React from "react";

interface FindingsProps {}

export const Findings: React.FC<FindingsProps> = ({}) => {
  const { data: allHostsData } = useAllHosts();
  const { data: allCredsData } = useAllCreds();
  return (
    <div className="container py-10">
      <Heading>Findings & Results</Heading>
      <div className="h-10" />
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-12 w-full  mt-10"> 
      {JSON.stringify(allHostsData)}
      {JSON.stringify(allCredsData)}
        
        </div>
    </div>
  );
};

export default Findings;
