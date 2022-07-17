import ddddocr

ocr = ddddocr.DdddOcr(det=False, ocr=True, show_ad=False)
with open('tmp/test.png', 'rb') as f:
# with open('tmp/verify_code.png', 'rb') as f:
    image_bytes = f.read()

res = ocr.classification(image_bytes)
print(res)