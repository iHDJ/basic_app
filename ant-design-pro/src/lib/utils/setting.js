export const { gon = {} } = window;

export const isProduction = gon.is_production || false;
export const htmlBaseTitle = gon.html_base_title || "Agideo";
export const copyrightDesc = gon.copyright || "2019 Agideo.com";
export const defaultLocale = gon.default_locale;
export const directUploadUrl =
  gon.direct_upload_url || "/api/direct_uploads";
export const sentryDsn = gon.sentry_js_dsn;
export const requestHeaders = gon.request_headers || {};
export const csrfToken = getMetaContent("csrf-token");
export const frontendVersion = getMetaContent("frontend-version");

function getMetaContent(name) {
  const meta = document.head.querySelector(`meta[name='${name}']`);

  if (!meta) {
    return null;
  }

  return meta.getAttribute("content");
}
