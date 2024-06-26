// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package datepicker

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

func tpl(props *Props) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{\n      datePickerOpen: false,\n      datePickerValue: &#39;&#39;,\n      datePickerFormat: &#39;M d, Y&#39;,\n      datePickerMonth: &#39;&#39;,\n      datePickerYear: &#39;&#39;,\n      datePickerDay: &#39;&#39;,\n      datePickerDaysInMonth: [],\n      datePickerBlankDaysInMonth: [],\n      datePickerMonthNames: [&#39;January&#39;, &#39;February&#39;, &#39;March&#39;, &#39;April&#39;, &#39;May&#39;, &#39;June&#39;, &#39;July&#39;, &#39;August&#39;, &#39;September&#39;, &#39;October&#39;, &#39;November&#39;, &#39;December&#39;],\n      datePickerDays: [&#39;Sun&#39;, &#39;Mon&#39;, &#39;Tue&#39;, &#39;Wed&#39;, &#39;Thu&#39;, &#39;Fri&#39;, &#39;Sat&#39;],\n      datePickerDayClicked(day) {\n        let selectedDate = new Date(this.datePickerYear, this.datePickerMonth, day);\n        this.datePickerDay = day;\n        this.datePickerValue = this.datePickerFormatDate(selectedDate);\n        this.datePickerIsSelectedDate(day);\n        this.datePickerOpen = false;\n      },\n      datePickerPreviousMonth(){\n        if (this.datePickerMonth == 0) { \n            this.datePickerYear--; \n            this.datePickerMonth = 12; \n        } \n        this.datePickerMonth--;\n        this.datePickerCalculateDays();\n      },\n      datePickerNextMonth(){\n        if (this.datePickerMonth == 11) { \n            this.datePickerMonth = 0; \n            this.datePickerYear++; \n        } else { \n            this.datePickerMonth++; \n        }\n        this.datePickerCalculateDays();\n      },\n      datePickerIsSelectedDate(day) {\n        const d = new Date(this.datePickerYear, this.datePickerMonth, day);\n        return this.datePickerValue === this.datePickerFormatDate(d) ? true : false;\n      },\n      datePickerIsToday(day) {\n        const today = new Date();\n        const d = new Date(this.datePickerYear, this.datePickerMonth, day);\n        return today.toDateString() === d.toDateString() ? true : false;\n      },\n      datePickerCalculateDays() {\n        let daysInMonth = new Date(this.datePickerYear, this.datePickerMonth + 1, 0).getDate();\n        // find where to start calendar day of week\n        let dayOfWeek = new Date(this.datePickerYear, this.datePickerMonth).getDay();\n        let blankdaysArray = [];\n        for (var i = 1; i &lt;= dayOfWeek; i++) {\n            blankdaysArray.push(i);\n        }\n        let daysArray = [];\n        for (var i = 1; i &lt;= daysInMonth; i++) {\n            daysArray.push(i);\n        }\n        this.datePickerBlankDaysInMonth = blankdaysArray;\n        this.datePickerDaysInMonth = daysArray;\n      },\n      datePickerFormatDate(date) {\n        let formattedDay = this.datePickerDays[date.getDay()];\n        let formattedDate = (&#39;0&#39; + date.getDate()).slice(-2); // appends 0 (zero) in single digit date\n        let formattedMonth = this.datePickerMonthNames[date.getMonth()];\n        let formattedMonthShortName = this.datePickerMonthNames[date.getMonth()].substring(0, 3);\n        let formattedMonthInNumber = (&#39;0&#39; + (parseInt(date.getMonth()) + 1)).slice(-2);\n        let formattedYear = date.getFullYear();\n\n        if (this.datePickerFormat === &#39;M d, Y&#39;) {\n          return `${formattedMonthShortName} ${formattedDate}, ${formattedYear}`;\n        }\n        if (this.datePickerFormat === &#39;MM-DD-YYYY&#39;) {\n          return `${formattedMonthInNumber}-${formattedDate}-${formattedYear}`;\n        }\n        if (this.datePickerFormat === &#39;DD-MM-YYYY&#39;) {\n          return `${formattedDate}-${formattedMonthInNumber}-${formattedYear}`;\n        }\n        if (this.datePickerFormat === &#39;YYYY-MM-DD&#39;) {\n          return `${formattedYear}-${formattedMonthInNumber}-${formattedDate}`;\n        }\n        if (this.datePickerFormat === &#39;D d M, Y&#39;) {\n          return `${formattedDay} ${formattedDate} ${formattedMonthShortName} ${formattedYear}`;\n        }\n        \n        return `${formattedMonth} ${formattedDate}, ${formattedYear}`;\n      },\n    }\" x-init=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf(`
        currentDate = new Date("%s");
        if (datePickerValue) {
            currentDate = new Date(Date.parse(datePickerValue));
        }
        datePickerMonth = currentDate.getMonth();
        datePickerYear = currentDate.getFullYear();
        datePickerDay = currentDate.getDay();
        datePickerValue = datePickerFormatDate( currentDate );
        datePickerCalculateDays();
        `, props.StartDate))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/datepicker/datepicker.templ`, Line: 103, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-cloak id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/datepicker/datepicker.templ`, Line: 105, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Label != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label for=\"datepicker\" class=\"block mb-1 text-sm font-medium text-neutral-500\">Select Date</label>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input x-ref=\"datePickerInput\" type=\"text\" @click=\"datePickerOpen=!datePickerOpen\" x-model=\"datePickerValue\" x-on:keydown.escape=\"datePickerOpen=false\" class=\"flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md text-neutral-600 border-neutral-300 ring-offset-background placeholder:text-neutral-400 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50\" placeholder=\"Select date\" readonly><div @click=\"datePickerOpen=!datePickerOpen; if(datePickerOpen){ $refs.datePickerInput.focus() }\" class=\"absolute top-0 right-0 px-3 py-2 cursor-pointer text-neutral-400 hover:text-neutral-500\"><svg class=\"w-6 h-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z\"></path></svg></div><div x-show=\"datePickerOpen\" x-transition @click.away=\"datePickerOpen = false\" class=\"absolute top-0 left-0 max-w-lg p-4 mt-12 antialiased bg-white border rounded-lg shadow w-[17rem] border-neutral-200/70\"><div class=\"flex items-center justify-between mb-2\"><div><span x-text=\"datePickerMonthNames[datePickerMonth]\" class=\"text-lg font-bold text-gray-800\"></span> <span x-text=\"datePickerYear\" class=\"ml-1 text-lg font-normal text-gray-600\"></span></div><div><button @click=\"datePickerPreviousMonth()\" type=\"button\" class=\"inline-flex p-1 transition duration-100 ease-in-out rounded-full cursor-pointer focus:outline-none focus:shadow-outline hover:bg-gray-100\"><svg class=\"inline-flex w-6 h-6 text-gray-400\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 19l-7-7 7-7\"></path></svg></button> <button @click=\"datePickerNextMonth()\" type=\"button\" class=\"inline-flex p-1 transition duration-100 ease-in-out rounded-full cursor-pointer focus:outline-none focus:shadow-outline hover:bg-gray-100\"><svg class=\"inline-flex w-6 h-6 text-gray-400\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 5l7 7-7 7\"></path></svg></button></div></div><div class=\"grid grid-cols-7 mb-3\"><template x-for=\"(day, index) in datePickerDays\" :key=\"index\"><div class=\"px-0.5\"><div x-text=\"day\" class=\"text-xs font-medium text-center text-gray-800\"></div></div></template></div><div class=\"grid grid-cols-7\"><template x-for=\"blankDay in datePickerBlankDaysInMonth\"><div class=\"p-1 text-sm text-center border border-transparent\"></div></template><template x-for=\"(day, dayIndex) in datePickerDaysInMonth\" :key=\"dayIndex\"><div class=\"px-0.5 mb-1 aspect-square\"><div x-text=\"day\" @click=\"datePickerDayClicked(day)\" :class=\"{\n                                        &#39;bg-neutral-200&#39;: datePickerIsToday(day) == true, \n                                        &#39;text-gray-600 hover:bg-neutral-200&#39;: datePickerIsToday(day) == false &amp;&amp; datePickerIsSelectedDate(day) == false,\n                                        &#39;bg-neutral-800 text-white hover:bg-opacity-75&#39;: datePickerIsSelectedDate(day) == true\n                                    }\" class=\"flex items-center justify-center text-sm leading-none text-center rounded-full cursor-pointer h-7 w-7\"></div></div></template></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
