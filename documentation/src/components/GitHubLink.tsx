import React from "react";

export enum GitHubType {
  DISCUSSION = "discussions",
  ISSUE = "issues",
  PULL_REQUEST = "pull",
}

type GitHubLinkProps = {
  id: number;
  type: GitHubType;
};

export default function GitHubLink({ id, type }: GitHubLinkProps) {
  return <a href={`https://github.com/kkrypt0nn/centauri/${type}/${id}`}>GH-{id}</a>;
}
