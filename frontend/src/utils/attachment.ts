import imageCompression from "browser-image-compression";

const COMPRESS_OPTIONS = Object.freeze({
  maxSizeMB: 0.5,
});

export const compressFileAndConvertToUint8Array = (
  file: File
): Promise<Uint8Array> => {
  return imageCompression(file, COMPRESS_OPTIONS).then((compressed) =>
    new Response(compressed)
      .arrayBuffer()
      .then((arrayBuffer) => new Uint8Array(arrayBuffer))
  );
};
