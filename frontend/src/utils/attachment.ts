export const blobToUint8Array = (blob: Blob): Promise<Uint8Array> => {
  return new Response(blob)
    .arrayBuffer()
    .then((arrayBuffer) => new Uint8Array(arrayBuffer));
};
