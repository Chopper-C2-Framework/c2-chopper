import React from "react";

interface HeadingProps {
  children: React.ReactNode;
}

export const Heading: React.FC<HeadingProps> = ({ children }) => {
  return <h1 className="text-5xl font-bold text-primary">{children}</h1>;
};

export default Heading;
