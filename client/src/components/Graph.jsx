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
                {props.fields.map((field, i) => <p key={i}>{field}</p>)}
            </div>
        );
    }
}

const Graph = (props) => {
    const graphRef = useRef(null);
    const [topState, setTopState] = useState(0);
    const [leftState, setLeftState] = useState(0);
    const [fieldsState, setFieldsState] = useState([]);

    const maxVal = data => {
        let maxX = 0;
        let maxY = 0;
        data.forEach(e => {
            maxX = Math.max(maxX, Math.abs(e.pos));
            maxY = Math.max(maxY, Math.abs(e.energy));
        });
        
        return [maxX, maxY];
    }

    useEffect(() => {
        console.log(props.data);
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
                .style("fill", "orange")
                .attr("r", 5)
                .attr("cx", d => {return x(d.pos)})
                .attr("cy", d => {return y(d.energy)})
                .on("mouseover", (d, i) => {
                    setLeftState(d.x);
                    setTopState(d.y);
                    setFieldsState([
                        `Positivity: ${d.target.__data__.pos}`,
                        `Energy: ${d.target.__data__.energy}`,
                        `Title: ${d.target.__data__.name}`,
                        `Artist: ${d.target.__data__.artist}`
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
            
            const xTextOffset = (height / 2) + 30;
            const yTextOffset = (width / 2) + 20;
            svg.append("text").text("Positivity")
                .attr("x", 40)
                .attr("y", xTextOffset)
                .style("text-anchor", "middle")
                .style("color", "white");
        
            svg.append("text").text("Energy")
                .attr("x", yTextOffset)
                .attr("y", 20)
                .style("text-anchor", "middle")
                .style("color", "white");
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