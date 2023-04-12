import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Layout from "@theme/Layout";
import React from "react";

export default function Home() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Centauri is a basic and simple Discord API wrapper written in Go."
    >
      <main>
        <div class="container padding-top--md padding-bottom--lg text--center">
          <div class="markdown">
            <h1>{siteConfig.title}</h1>
            <p>
              Centauri is a basic and simple Discord API wrapper written in Go.
              <br />
              <br />
              If you like this project, do not hesitate to <strong>
                star
              </strong>{" "}
              the repository{" "}
              <a href="https://github.com/kkrypt0nn/centauri" target="_blank">
                <img
                  alt="GitHub Repo stars"
                  src="https://img.shields.io/github/stars/kkrypt0nn/Centauri?style=social"
                />
              </a>{" "}
              and check the <a href="/donations">donations page</a>.
            </p>
          </div>
        </div>
      </main>
    </Layout>
  );
}
