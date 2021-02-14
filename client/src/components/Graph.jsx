import {React, useRef, useEffect, useState} from 'react';
import * as d3 from "d3";
import "../styles/Graph.css";

const ToolTip = (props) => {
    if (props.left === 0 || props.top === 0) {
        return (null);
    }
    else {
        return (
            <div className="tooltip" style={{ left: `${props.left + 15}px`, top: `${props.top + 15}px` }}>
                <p id="tooltip-header">{props.fields[0]}</p>
                <p>{props.fields[1]}</p>
                <p>{props.fields[2]}</p>
            </div>
        );
    }
}

const Graph = (props) => {
    const graphRef = useRef(null);
    const [topState, setTopState] = useState(0);
    const [leftState, setLeftState] = useState(0);
    const [fieldsState, setFieldsState] = useState([]);

    useEffect(() => {
        if (props.data && graphRef.current) {
            d3.selectAll("svg").remove();
            const width = 700;
            const height = 700;
            let svg = d3.select(graphRef.current)
                .append("svg")
                .attr("width", width)
                .attr("height", height)
            const data = props.data;
            let x = d3.scaleLinear().domain([0,1]).range([50, width - 50]);
            let y = d3.scaleLinear().domain([0,1]).range([height - 50, 50]);

            svg.selectAll("circle")
                .data(data).enter().append("circle")
                .style("fill", "#1db954")
                .attr("r", 5)
                .attr("cx", d => {return x(d.pos)})
                .attr("cy", d => {return y(d.energy)})
                .on("mouseover", (d) => {
                    setLeftState(d.x);
                    setTopState(d.y);
                    setFieldsState([
                        `${d.target.__data__.artist} - ${d.target.__data__.name}`,
                        `Positivity: ${d.target.__data__.pos}`,
                        `Energy: ${d.target.__data__.energy}`,
                    ]);
                })
                .on("mouseout", d => {
                    setLeftState(0);
                    setTopState(0);
                    setFieldsState([]);
                })
                .exit().remove();
            
            const xAxis = d3.axisBottom().tickValues([]).tickSizeOuter(0).scale(x);
            const yAxis = d3.axisLeft().tickValues([]).tickSizeOuter(0).scale(y);
            const xTranslate = height/2;
            const yTranslate = width/2;
            svg.append("g").attr("transform", "translate(0," +  xTranslate + ")").call(xAxis);
            svg.append("g").attr("transform", "translate(" + yTranslate + ", 0)").call(yAxis);
            d3.json("data.json")
            
            const xTextOffset = (height / 2) - 15;
            const yTextOffset = (width / 2);
            svg.append("text").text("Positivity +")
                .attr("x", width-50)
                .attr("y", xTextOffset)
                .style("text-anchor", "middle")
                .style("fill", "white")
                .style("font-weight", "700");
        
            svg.append("text").text("Energy +")
                .attr("x", yTextOffset)
                .attr("y", 20)
                .style("text-anchor", "middle")
                .style("fill", "white")
                .style("font-weight", "700");
        }
        
    }, [props.data]);

    return (
        <div>
            <div ref={graphRef} className="main"/>
            <ToolTip left={leftState} top={topState} fields={fieldsState}/>
        </div>
    );
}

export default Graph;