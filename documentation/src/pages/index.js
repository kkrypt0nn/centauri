import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Admonition from "@theme/Admonition";
import Layout from "@theme/Layout";
import React from "react";

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Hello from Centauri`}
      description="Centauri is a basic and simple Discord API wrapper written in Go."
    >
      <main>
        <div class="container padding-top--md padding-bottom--lg">
          <div class="markdown">
            <h1>Centauri</h1>
            <p className="spaced">
              <a href="https://pkg.go.dev/github.com/kkrypt0nn/centauri">
                <img
                  src="https://pkg.go.dev/badge/github.com/kkrypt0nn/centauri.svg"
                  alt="Go Reference Badge"
                />
              </a>{" "}
              <a href="https://github.com/kkrypt0nn/centauri/actions">
                <img
                  src="https://github.com/kkrypt0nn/centauri/actions/workflows/ci.yml/badge.svg"
                  alt="CI Badge"
                />
              </a>{" "}
              <a href="https://discord.gg/feA6ZGRgpw">
                <img
                  src="https://img.shields.io/discord/1095412499416891482?logo=discord"
                  alt="Discord Server Badge"
                />
              </a>{" "}
              <a href="https://github.com/kkrypt0nn/centauri/commits/main">
                <img
                  src="https://img.shields.io/github/last-commit/kkrypt0nn/centauri"
                  alt="Last Commit Badge"
                />
              </a>{" "}
              <a href="https://conventionalcommits.org/en/v1.0.0/">
                <img
                  src="https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&amp;logoColor=white"
                  alt="Conventional Commits Badge"
                />
              </a>
            </p>
            <Admonition type="note" title="Note">
              <p>
                Due to my military service from <strong>July 2023</strong> until{" "}
                <strong>April 2024</strong> I will not be able to work on this
                project a lot, most likely only during weekends if I have time
                and motivation ^-^
              </p>
            </Admonition>
            <Admonition type="caution" title="Warning">
              <p>
                This is a WIP library that may be <strong>very unstable</strong>{" "}
                and eventually <strong>not fully optimised</strong>, use at your
                own care! This page will also be reworked.
              </p>
            </Admonition>
            <p>
              Centauri is a basic and simple Discord API wrapper written in Go.
            </p>
            <h2 id="features">Features</h2>
            <p>
              The plan for Centauri would be for it to cover Discord's APIs in
              their entirety. A small list of features that will be in Centauri
              are:
            </p>
            <ul>
              <li>REST API</li>
              <li>Gateway API</li>
              <li>Interactions (Buttons, Modals, etc.)</li>
              <li>Tasks</li>
              <li>ANSI Colors</li>
            </ul>
            <h2 id="examples">Examples</h2>
            <p>
              Examples are available{" "}
              <a href="https://github.com/kkrypt0nn/centauri/tree/main/_examples">
                here
              </a>
              .
            </p>
            <h2 id="contributing">Contributing</h2>
            <p>
              Before the version <strong>1.0.0</strong> of the project,
              contributions are mostly closed as I&#39;d like to lay down the
              foundations of the library. Afterward contributions will be more
              than welcome if following the{" "}
              <a href="https://centauri.krypton.ninja/community/contributing-guidelines/">
                Contributing Guidelines
              </a>
              .
            </p>
            <h2 id="license">License</h2>
            <p>
              This library was made with 💜 by Krypton and is under the{" "}
              <a href="https://github.com/kkrypt0nn/centauri/blob/main/LICENSE.md">
                GPL v3
              </a>{" "}
              license.
            </p>
          </div>
        </div>
      </main>
    </Layout>
  );
}
