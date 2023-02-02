package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher 
		label: image-converter,  

        // application identifier name
        name: "org.rednib.imageformatter",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width:  "1:1",
        height: "500px",
        title:  "Image Converter",
        content: {
            type: main.Home,
        },
        background: #ffffff,
    },
}
`
