function editTableCountries() {
    let CountriesName;
    let Flag;
    let ReligionID;
    let LanguagesID;
    let GovernmentFormID;
    let TerritorySizeID;

    let data;
    let TableName = "Countries";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        CountriesName = document.getElementById("CountriesName").value
        Flag = document.getElementById("Flag").value
        ReligionID = document.getElementById("ReligionID").value
        LanguagesID = document.getElementById("LanguagesID").value
        GovernmentFormID = document.getElementById("GovernmentFormID").value
        TerritorySizeID = document.getElementById("TerritorySizeID").value

        data = "TableName=" + TableName + "&CountriesName=" + encodeURIComponent(CountriesName) + "&Flag=" +
            encodeURIComponent(Flag) + "&ReligionID=" + encodeURIComponent(ReligionID) +
            "&LanguagesID=" + encodeURIComponent(LanguagesID) + "&GovernmentFormID=" +
            encodeURIComponent(GovernmentFormID) + "&TerritorySizeID=" + encodeURIComponent(TerritorySizeID)
        xhr.open("POST", '/table/insert', true);
    }

    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableLanguages() {
    let Language;

    let data;
    let TableName = "Languages";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        Language = document.getElementById("Language").value

        data = "TableName=" + TableName + "&Language=" + encodeURIComponent(Language)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableGovernmentForms() {
    let Form;

    let data;
    let TableName = "GovernmentForms";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        Form = document.getElementById("Form").value

        data = "TableName=" + TableName + "&Form=" + encodeURIComponent(Form)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableTerritorySizes() {
    let Type;

    let data;
    let TableName = "TerritorySizes";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        Type = document.getElementById("Type").value

        data = "TableName=" + TableName + "&Type=" + encodeURIComponent(Type)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}
function editTableReligions() {
    let Title;

    let data;
    let TableName = "Religions";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        Title = document.getElementById("Title").value

        data = "TableName=" + TableName + "&Title=" + encodeURIComponent(Title)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableAssociations() {
    let Title;

    let data;
    let TableName = "Associations";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        Title = document.getElementById("Title").value

        data = "TableName=" + TableName + "&Title=" + encodeURIComponent(Title)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableAssociationsOfCountries() {
    let CountriesID;
    let AssociationsID;

    let data;
    let TableName = "AssociationsOfCountries";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        CountriesID = document.getElementById("CountriesID").value
        AssociationsID = document.getElementById("AssociationsID").value

        data = "TableName=" + TableName + "&CountriesID=" + encodeURIComponent(CountriesID) + "&AssociationsID=" +
            encodeURIComponent(AssociationsID)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableContinents() {
    let Name;

    let data;
    let TableName = "Continents";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        Name = document.getElementById("Name").value

        data = "TableName=" + TableName + "&Name=" + encodeURIComponent(Name)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

function editTableContinentsOfCountries() {
    let CountriesID;
    let ContinentsID;

    let data;
    let TableName = "СontinentsOfCountries";
    const xhr = new XMLHttpRequest();
    if (document.getElementById("specificSizeSelect").options.selectedIndex === 1) {
        CountriesID = document.getElementById("CountriesID").value
        ContinentsID = document.getElementById("ContinentsID").value

        data = "TableName=" + TableName + "&CountriesID=" + encodeURIComponent(CountriesID) + "&ContinentsID=" +
            encodeURIComponent(ContinentsID)
        xhr.open("POST", '/table/insert', true);
    }
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send(data);

    window.location.reload();
}

