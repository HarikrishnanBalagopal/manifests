# Kubeflow on IKS Website

Welcome to the GitHub repository for Kubeflow on IKS website!

The docs website is hosted at https://ibm.github.io/manifests/.

We use [Hugo](https://gohugo.io/) with the [google/docsy](https://github.com/google/docsy)
theme for styling and site structure, and [Netlify](https://www.netlify.com/) to manage the deployment of the site.

## Local development

This section will show you how to develop the website locally, by running a local Hugo server.

### Recursively download the submodules (for docsy):

Check out the submodule(doscy) as the main hugo theme module:
```sh
git submodule update --init --recursive
```

### Install Hugo

Building and running the site locally requires a recent `extended` version of [Hugo](https://gohugo.io).
You can find out more about how to install Hugo for your environment in our
[Getting started](https://www.docsy.dev/docs/getting-started/#prerequisites-and-installation) guide.

In this project, the Docsy theme component is pulled in as a Hugo module, together with other module dependencies:

```bash
$ hugo mod graph
hugo: collected modules in 566 ms
hugo: collected modules in 578 ms
github.com/google/docsy-example github.com/google/docsy@v0.5.1-0.20221017155306-99eacb09ffb0
github.com/google/docsy-example github.com/google/docsy/dependencies@v0.5.1-0.20221014161617-be5da07ecff1
github.com/google/docsy/dependencies@v0.5.1-0.20221014161617-be5da07ecff1 github.com/twbs/bootstrap@v4.6.2+incompatible
github.com/google/docsy/dependencies@v0.5.1-0.20221014161617-be5da07ecff1 github.com/FortAwesome/Font-Awesome@v0.0.0-20220831210243-d3a7818c253f
```

### Run local hugo server

Follow the usual GitHub workflow of forking the repository on GitHub and then cloning your fork to your local machine.

1. **Fork** the [kubeflow/website repository](https://github.com/kubeflow/website) in the GitHub UI.

2. Clone your fork locally:

    ```bash
    git clone git@github.com:<your-github-username>/manifests.git
    cd manifests/website
    ```

3. Install Node packages:

    ```bash
    npm install
    ```

4. Start your local Hugo server:

    ```bash
    hugo server
    ```

5. You can access your website at [http://localhost:1313/manifests/](http://localhost:1313/manifests/)


### Build the static pages

Use the following command to build the static pages and publish as github page:

1. Generate static pages:
  ```
  HUGO_ENV="production" hugo --gc -d <destination folder>
  ```

2. Copy static pages to `/github-pages` directory
  Currently, the github action is set up to publish the static pages under `github-pages` in `github-pages` branch when there
  is any update in `github-pages` branch.

3. Add a commit for the static pages and push to `github-pages` branch.

### Useful docs

* [User guide for the Docsy theme](https://www.docsy.dev/docs/getting-started/)
* [Hugo installation guide](https://gohugo.io/getting-started/installing/)
* [Hugo basic usage](https://gohugo.io/getting-started/usage/)
* [Hugo site directory structure](https://gohugo.io/getting-started/directory-structure/)
* [hugo server reference](https://gohugo.io/commands/hugo_server/)

## Menu structure

The site theme has one Hugo menu (`main`), which defines the top navigation bar. You can find and adjust the definition
of the menu in the [site configuration file](https://github.com/kubeflow/website/blob/master/config.toml).

The left-hand navigation panel is defined by the directory structure under the [`docs` directory](https://github.com/kubeflow/website/tree/master/content/en/docs).

A `weight` property in the _front matter_ of each page determines the position of the page relative to the others in the same directory.
The lower the weight, the earlier the page appears in the section.

Here is an example `_index.md` file:

```md
+++
title = "Getting Started with Kubeflow"
description = "Overview"
weight = 1
+++
```

## Using Hugo shortcodes

Sometimes it's useful to define a snippet of information in one place and reuse it wherever we need it. 
For example, we want to be able to refer to the minimum version of various frameworks/libraries throughout the docs, 
without causing a maintenance nightmare.

For this purpose, we use Hugo's "shortcodes". 
Shortcodes are similar to Django variables. You define a shortcode in a file, then use a specific markup 
to invoke the shortcode in the docs. That markup is replaced by the content of the shortcode file when the page is built.

To create a shortcode:

1. Add an HTML file in the `manifests/website/layouts/shortcodes/` directory. 
   The file name must be short and meaningful, as it determines the shortcode you and others use in the docs.

2. For the file content, add the text and HTML markup that should replace the shortcode markup when the web page is built.

To use a shortcode in a document, wrap the name of the shortcode in braces and percent signs like this:

  ```
  {{% shortcode-name %}}
  ```

The shortcode name is the file name minus the `.html` file extension.

**Example:** The following shortcode defines the minimum required version of Kubernetes:

* File name of the shortcode:

  ```
  kubernetes-min-version.html
  ```

* Content of the shortcode:

  ```
  1.8
  ```

* Usage in a document:

  ```
  You need Kubernetes version {{% kubernetes-min-version %}} or later.
  ```

Useful Hugo docs:

* [Shortcode templates](https://gohugo.io/templates/shortcode-templates/)
* [Shortcodes](https://gohugo.io/content-management/shortcodes/)

