// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "Centauri",
  tagline: "Centauri is a basic and simple Discord API wrapper written in Go.",
  favicon: "img/favicon.ico",

  url: "https://centauri.krypton.ninja",
  baseUrl: "/",

  organizationName: "kkrypt0nn",
  projectName: "centauri",
  trailingSlash: true,

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve("./sidebars.js"),
          editUrl:
            "https://github.com/kkrypt0nn/centauri/tree/main/documentation/",
        },
        blog: {
          showReadingTime: true,
        },
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      }),
    ],
    [
      "@docusaurus/plugin-sitemap",
      {
        sitemap: {
          changefreq: "weekly",
          priority: 0.5,
          filename: "sitemap.xml",
        },
      },
    ],
  ],

  plugins: [
    [
      "@docusaurus/plugin-content-docs",
      /** @type {import('@docusaurus/plugin-content-docs').Options} */
      ({
        id: "community",
        path: "community",
        routeBasePath: "community",
        sidebarPath: require.resolve("./sidebarsCommunity.js"),
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      metadata: [
        {
          name: "keywords",
          content:
            "centauri, discord, discord api, discord api wrapper, api wrapper, discord library, go library, golang library discord, discord go library",
        },
        {
          name: "theme-color",
          content: "#11111b",
        },
      ],
      image: "img/centauri-banner.png",
      colorMode: {
        defaultMode: "dark",
        disableSwitch: false,
        respectPrefersColorScheme: true,
      },
      announcementBar: {
        content:
          "⚠️ Centauri is currently in <b>alpha version</b> and under ongoing devlopment, meaning it may not be stable. This website is a work in progress as well!",
        backgroundColor: "var(--ifm-navbar-background-color)",
        textColor: "var(--ifm-font-color-base)",
        isCloseable: true,
      },
      navbar: {
        title: "Centauri",
        logo: {
          alt: "Centauri Logo",
          src: "img/centauri.png",
        },
        items: [
          {
            type: "doc",
            docId: "installation/README",
            position: "left",
            label: "Documentation",
          },
          {
            to: "/blog",
            position: "left",
            label: "Blog",
          },
          {
            to: "/community/code-of-conduct",
            label: "Community",
            position: "left",
          },
          {
            to: "/intents-calculator",
            label: "Intents Calculator",
            position: "left",
          },
          {
            href: "https://github.com/kkrypt0nn/centauri",
            label: "GitHub",
            position: "right",
          },
          {
            href: "https://discord.gg/feA6ZGRgpw",
            label: "Discord",
            position: "right",
          },
        ],
      },
      footer: {
        style: "dark",
        links: [
          {
            title: "Docs",
            items: [
              {
                label: "Documentation",
                to: "/docs/installation/",
              },
              {
                label: "Blog",
                to: "/blog/",
              },
              {
                label: "Community",
                to: "/community/code-of-conduct",
              },
              {
                label: "Intents Calculator",
                to: "/intents-calculator",
              },
            ],
          },
          {
            title: "Community",
            items: [
              {
                label: "Code of Conduct",
                to: "/community/code-of-conduct",
              },
              {
                label: "Contributing Guidelines",
                to: "/community/contributing-guidelines",
              },
              {
                label: "Security Policy",
                to: "/community/security-policy",
              },
              {
                label: "Discord",
                href: "https://discord.gg/feA6ZGRgpw",
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Centauri (Made by Krypton)`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
