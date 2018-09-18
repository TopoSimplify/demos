const mdpdf = require('mdpdf');
const path = require('path');

let options = {
    source: path.join(__dirname, 'README.md'),
    destination: "../dist/constdp/README.pdf",
    ghStyle : true,
    pdf: {
        format: 'A4',
        orientation: 'portrait'
    }
};

mdpdf.convert(options).then((pdfPath) => {}).catch((err) => {console.error(err);});